resource "aws_vpc" "angelowl" {
  cidr_block           = "10.10.10.0/23"
  enable_dns_support   = true
  enable_dns_hostnames = true

  tags = {
    Name = "angelowl-vpc"
  }
}

resource "aws_internet_gateway" "angelowl" {
  vpc_id = aws_vpc.angelowl.id

  tags = {
    Name = "angelowl-igw"
  }
}

# Private subnets and their routing tables

resource "aws_subnet" "angelowl_private_a" {
  vpc_id            = aws_vpc.angelowl.id
  cidr_block        = "10.10.10.0/26"
  availability_zone = "ap-southeast-1a"

  tags = {
    "kubernetes.io/cluster/angelowl-eks-cluster" = "shared"
    "kubernetes.io/role/internal-elb" = 1
  }
}

resource "aws_subnet" "angelowl_private_b" {
  vpc_id            = aws_vpc.angelowl.id
  cidr_block        = "10.10.10.64/26"
  availability_zone = "ap-southeast-1b"

  tags = {
    "kubernetes.io/cluster/angelowl-eks-cluster" = "shared"
    "kubernetes.io/role/internal-elb" = 1
  }
}

resource "aws_subnet" "angelowl_private_c" {
  vpc_id            = aws_vpc.angelowl.id
  cidr_block        = "10.10.10.128/26"
  availability_zone = "ap-southeast-1c"

  tags = {
    "kubernetes.io/cluster/angelowl-eks-cluster" = "shared"
    "kubernetes.io/role/internal-elb" = 1
  }
}

resource "aws_route_table" "angelowl_private_default" {
  vpc_id = aws_vpc.angelowl.id

  route {
    cidr_block     = "0.0.0.0/0"
    nat_gateway_id = aws_nat_gateway.angelowl_nat.id
  }

  tags = {
    Name = "angelowl_private_default_rtb"
  }
}

resource "aws_route_table_association" "angelowl_private_a" {
  subnet_id      = aws_subnet.angelowl_private_a.id
  route_table_id = aws_route_table.angelowl_private_default.id
}

resource "aws_route_table_association" "angelowl_private_b" {
  subnet_id      = aws_subnet.angelowl_private_b.id
  route_table_id = aws_route_table.angelowl_private_default.id
}

resource "aws_route_table_association" "angelowl_private_c" {
  subnet_id      = aws_subnet.angelowl_private_c.id
  route_table_id = aws_route_table.angelowl_private_default.id
}

# Public subnets and their routing tables

resource "aws_subnet" "angelowl_public_a" {
  vpc_id            = aws_vpc.angelowl.id
  cidr_block        = "10.10.11.0/26"
  availability_zone = "ap-southeast-1a"
}

resource "aws_subnet" "angelowl_public_b" {
  vpc_id            = aws_vpc.angelowl.id
  cidr_block        = "10.10.11.64/26"
  availability_zone = "ap-southeast-1b"
}

resource "aws_subnet" "angelowl_public_c" {
  vpc_id            = aws_vpc.angelowl.id
  cidr_block        = "10.10.11.128/26"
  availability_zone = "ap-southeast-1c"
}

resource "aws_route_table" "angelowl_public_default" {
  vpc_id = aws_vpc.angelowl.id

  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.angelowl.id
  }

  tags = {
    Name = "angelowl_public_default_rtb"
  }
}

resource "aws_route_table_association" "angelowl_public_a" {
  route_table_id = aws_route_table.angelowl_public_default.id
  subnet_id      = aws_subnet.angelowl_public_a.id
}

resource "aws_route_table_association" "angelowl_public_b" {
  route_table_id = aws_route_table.angelowl_public_default.id
  subnet_id      = aws_subnet.angelowl_public_b.id
}

resource "aws_route_table_association" "angelowl_public_c" {
  route_table_id = aws_route_table.angelowl_public_default.id
  subnet_id      = aws_subnet.angelowl_public_c.id
}

# NAT gateway for internet access

resource "aws_nat_gateway" "angelowl_nat" {
  allocation_id = aws_eip.angelowl_nat.id
  subnet_id     = aws_subnet.angelowl_public_a.id
  private_ip    = "10.10.11.5"

  tags = {
    Name = "angelowl-natgw"
  }

  depends_on = [
    aws_internet_gateway.angelowl
  ]
}

resource "aws_eip" "angelowl_nat" {
  vpc = true
}
