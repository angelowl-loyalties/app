data "aws_iam_policy_document" "eks_assume_role_policy" {
  statement {
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["eks.amazonaws.com"]
    }
  }
}

data "aws_iam_policy_document" "eks_pod_assume_role_policy" {
  statement {
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["eks-fargate-pods.amazonaws.com"]
    }
  }
}

data "aws_iam_policy_document" "eks_node_assume_role_policy" {
  statement {
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["ec2.amazonaws.com"]
    }
  }
}

data "aws_iam_policy_document" "eks_ebs_assume_role_policy" {
  statement {
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = [
        "ec2.amazonaws.com",
        "eks-fargate-pods.amazonaws.com",
        "eks.amazonaws.com"
      ]
    }
  }

  statement {
    actions = ["sts:AssumeRoleWithWebIdentity"]

    principals {
      type        = "Federated"
      identifiers = ["arn:aws:iam::276374573009:oidc-provider/oidc.eks.ap-southeast-1.amazonaws.com/id/49E37774B5EBCFA6A3D8329700DA0A05"]
    }

    condition {
      test     = "StringEquals"
      variable = "oidc.eks.ap-southeast-1.amazonaws.com/id/49E37774B5EBCFA6A3D8329700DA0A05:aud"

      values = [
        "sts.amazonaws.com",
      ]
    }

    condition {
      test     = "StringEquals"
      variable = "oidc.eks.ap-southeast-1.amazonaws.com/id/49E37774B5EBCFA6A3D8329700DA0A05:sub"

      values = [
        "system:serviceaccount:kube-system:ebs-csi-controller-sa",
      ]
    }
  }
}

# Create the policy to allow ECR access
data "aws_iam_policy_document" "ecr_pull_policy_document" {
  statement {
    effect = "Allow"
    actions = [
      "ecr:BatchCheckLayerAvailability",
      "ecr:BatchGetImage",
      "ecr:GetDownloadUrlForLayer",
      "ecr:GetAuthorizationToken"
    ]
    resources = [
      "*"
    ]
  }
}

resource "aws_iam_policy" "angelowl_ebs_policy" {
  name        = "angel-owl-eks-ebs-policy"
  description = "Policy to allow EKS to attach EBS volumes to nodes"
  policy      = "${file("iam_ebs_policy.json")}"
}

resource "aws_iam_policy" "angelowl_alb_policy" {
  name        = "angel-owl-eks-alb-policy"
  description = "Policy to allow EKS to create load balancers for the cluster"
  policy      = "${file("iam_alb_policy.json")}"
}

resource "aws_iam_policy" "ecr_pull_policy" {
  name        = "angel-owl-eks-ecr-policy"
  description = "Policy to allow EKS to pull images from ECR"
  policy      = data.aws_iam_policy_document.ecr_pull_policy_document.json
}

resource "aws_iam_role" "eks-role" {
  name               = "angel-owl-eks-role"
  path               = "/"
  assume_role_policy = data.aws_iam_policy_document.eks_assume_role_policy.json
  managed_policy_arns = [
    aws_iam_policy.ecr_pull_policy.arn,
    "arn:aws:iam::aws:policy/AmazonEKSClusterPolicy"
  ]
}

resource "aws_iam_role" "eks-pod-role" {
  name                = "angel-owl-eks-pod-role"
  path                = "/"
  assume_role_policy  = data.aws_iam_policy_document.eks_pod_assume_role_policy.json
  managed_policy_arns = [aws_iam_policy.ecr_pull_policy.arn]
}

resource "aws_iam_role" "eks-node-role" {
  name               = "angel-owl-eks-node-role"
  path               = "/"
  assume_role_policy = data.aws_iam_policy_document.eks_node_assume_role_policy.json
  managed_policy_arns = [
    "arn:aws:iam::aws:policy/AmazonEKSWorkerNodePolicy",
    "arn:aws:iam::aws:policy/AmazonEC2ContainerRegistryReadOnly",
    "arn:aws:iam::aws:policy/AmazonEKS_CNI_Policy",
  ]
}

resource "aws_iam_role" "eks_ebs_role" {
  name               = "angel-owl-eks-ebs-role"
  path               = "/"
  assume_role_policy = data.aws_iam_policy_document.eks_ebs_assume_role_policy.json
  managed_policy_arns = [
    aws_iam_policy.angelowl_ebs_policy.arn
  ]
}

output "eks_alb_policy_arn" {
  value = aws_iam_policy.angelowl_alb_policy.arn
}