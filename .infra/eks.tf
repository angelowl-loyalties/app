resource "aws_eks_cluster" "angelowl_eks_cluster" {
  name     = "angelowl-eks-cluster"
  role_arn = aws_iam_role.eks-role.arn

  enabled_cluster_log_types = [
    "api",
    "audit",
    "authenticator",
    "controllerManager",
    "scheduler",
  ]

  vpc_config {
    endpoint_private_access = true
    endpoint_public_access  = true

    subnet_ids = [
      aws_subnet.angelowl_private_a.id,
      aws_subnet.angelowl_private_b.id,
      aws_subnet.angelowl_private_c.id,
    ]

    security_group_ids = [
      aws_security_group.angelowl_outbound.id,
      aws_security_group.angelowl_kubeservices.id,
      aws_security_group.angelowl_http_s_ingress.id,
    ]
  }

  encryption_config {
    resources = ["secrets"]

    provider {
      key_arn = aws_kms_key.angelowl_kms.arn
    }
  }
}

resource "aws_eks_node_group" "angelowl_eks_node_group" {
  cluster_name    = aws_eks_cluster.angelowl_eks_cluster.name
  node_group_name = "angelowl-eks-node-group"
  node_role_arn   = aws_iam_role.eks-node-role.arn

  subnet_ids = [
    aws_subnet.angelowl_private_a.id,
    aws_subnet.angelowl_private_b.id,
    aws_subnet.angelowl_private_c.id,
  ]

  scaling_config {
    desired_size = 3
    max_size     = 10
    min_size     = 3
  }

  instance_types = ["t3.small"]
  ami_type       = "AL2_x86_64"
  disk_size      = 40

  remote_access {
    ec2_ssh_key = aws_key_pair.angelowl_k3s.key_name
    source_security_group_ids = [
      aws_security_group.angelowl_outbound.id,
      aws_security_group.angelowl_ssh.id,
    ]
  }
}

resource "aws_eks_fargate_profile" "default_namespace" {
  fargate_profile_name = "angelowl-eks-fargate-default"

  cluster_name           = aws_eks_cluster.angelowl_eks_cluster.name
  pod_execution_role_arn = aws_iam_role.eks-pod-role.arn

  subnet_ids = [
    aws_subnet.angelowl_private_a.id,
    aws_subnet.angelowl_private_b.id,
    aws_subnet.angelowl_private_c.id,
  ]

  selector {
    namespace = "default"
  }
}

resource "aws_eks_fargate_profile" "coredns" {
  fargate_profile_name = "angelowl-eks-fargate-coredns"

  cluster_name           = aws_eks_cluster.angelowl_eks_cluster.name
  pod_execution_role_arn = aws_iam_role.eks-pod-role.arn

  subnet_ids = [
    aws_subnet.angelowl_private_a.id,
    aws_subnet.angelowl_private_b.id,
    aws_subnet.angelowl_private_c.id,
  ]

  selector {
    namespace = "kube-system"
    labels = {
      "k8s-app" = "kube-dns"
    }
  }
}

resource "aws_eks_addon" "kube-proxy" {
  addon_name        = "kube-proxy"
  cluster_name      = aws_eks_cluster.angelowl_eks_cluster.name
  resolve_conflicts = "OVERWRITE"
  addon_version     = "v1.24.9-eksbuild.1"
}

resource "aws_eks_addon" "vpc-cni" {
  addon_name        = "vpc-cni"
  cluster_name      = aws_eks_cluster.angelowl_eks_cluster.name
  resolve_conflicts = "OVERWRITE"
  addon_version     = "v1.12.2-eksbuild.1"
}

resource "aws_eks_addon" "coredns" {
  addon_name        = "coredns"
  cluster_name      = aws_eks_cluster.angelowl_eks_cluster.name
  resolve_conflicts = "OVERWRITE"
  addon_version     = "v1.9.3-eksbuild.2"
}

resource "aws_eks_addon" "ebs" {
  addon_name               = "aws-ebs-csi-driver"
  cluster_name             = aws_eks_cluster.angelowl_eks_cluster.name
  resolve_conflicts        = "OVERWRITE"
  addon_version            = "v1.16.0-eksbuild.1"
  service_account_role_arn = aws_iam_role.eks_ebs_role.arn
}

output "eks_endpoint" {
  value = aws_eks_cluster.angelowl_eks_cluster.endpoint
}

output "oidc_issuer" {
  value = aws_eks_cluster.angelowl_eks_cluster.identity.0.oidc.0.issuer
}
