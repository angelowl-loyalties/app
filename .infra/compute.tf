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
