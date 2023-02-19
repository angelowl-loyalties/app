data "aws_iam_policy_document" "eks_assume_role_policy" {
  statement {
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["eks.amazonaws.com"]
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

resource "aws_iam_policy" "ecr_pull_policy" {
  name        = "angel-owl-eks-ecr-policy"
  description = "Policy to allow EKS to pull images from ECR"
  policy      = data.aws_iam_policy_document.ecr_pull_policy_document.json
}

resource "aws_iam_role" "eks-role" {
  name                = "angel-owl-eks-role"
  path                = "/"
  assume_role_policy  = data.aws_iam_policy_document.eks_assume_role_policy.json
  managed_policy_arns = [aws_iam_policy.ecr_pull_policy.arn]
}
