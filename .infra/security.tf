resource "aws_security_group" "angelowl_ssh" {
  name        = "angelowl-ssh-access"
  description = "Allows inbound SSH traffic"
  vpc_id      = aws_vpc.angelowl.id

  ingress {
    description = "SSH"
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_security_group" "angelowl_outbound" {
  name        = "angelowl-outbound-access"
  description = "Allows outbound traffic"
  vpc_id      = aws_vpc.angelowl.id

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

