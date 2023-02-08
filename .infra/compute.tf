resource "aws_instance" "angelowl_vpn_instance" {
  ami           = "ami-0b7e55206a0a22afc"
  instance_type = "t3.nano"
  key_name      = aws_key_pair.angelowl_vpn.key_name
  subnet_id     = aws_subnet.angelowl_public_a.id


  vpc_security_group_ids = [
    aws_security_group.angelowl_ssh.id,
    aws_security_group.angelowl_openvpn.id,
    aws_security_group.angelowl_outbound.id
  ]

  root_block_device {
    delete_on_termination = false
    volume_size           = 20
    encrypted             = false
    volume_type           = "gp2"

    tags = {
      Name = "angelowl-vpn-ec2-storage"
    }
  }

  tags = {
    Name = "angelowl-vpn"
  }
}

resource "aws_eip" "angelowl_vpn_eip" {
  vpc      = true
  instance = aws_instance.angelowl_vpn_instance.id
}

output "angelowl_vpn_public_ip" {
  value = aws_eip.angelowl_vpn_eip.public_ip
}

# Mini K3s cluster for development purposes

resource "aws_instance" "angelowl_k3s_instance" {
  ami           = "ami-0b7e55206a0a22afc"
  instance_type = "t3.small"
  key_name      = aws_key_pair.angelowl_k3s.key_name
  subnet_id     = aws_subnet.angelowl_private_a.id
  private_ip    = "10.10.10.10"

  vpc_security_group_ids = [
    aws_security_group.angelowl_ssh.id,
    aws_security_group.angelowl_outbound.id,
    aws_security_group.angelowl_kubeservices.id
  ]

  root_block_device {
    delete_on_termination = false
    volume_size           = 50
    encrypted             = false
    volume_type           = "gp2"

    tags = {
      Name = "angelowl-k3s-ec2-storage"
    }
  }

  tags = {
    Name = "angelowl-k3s"
  }
}

output "angelowl_k3s_private_ip" {
  value = aws_instance.angelowl_k3s_instance.private_ip
}
