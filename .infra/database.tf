resource "aws_db_subnet_group" "database" {
  name = "angelowl-rds-subnet-group"

  subnet_ids = [
    aws_subnet.angelowl_private_a.id,
    aws_subnet.angelowl_private_b.id,
    aws_subnet.angelowl_private_c.id
  ]
}

resource "aws_db_instance" "campaignex_db" {
  allocated_storage = 20
  storage_type      = "gp3"

  db_subnet_group_name = aws_db_subnet_group.database.name
  vpc_security_group_ids = [
    aws_security_group.angelowl_postgres.id
  ]

  storage_encrypted = true
  kms_key_id        = "arn:aws:kms:ap-southeast-1:276374573009:key/5e9b5264-8d3d-4496-8682-7268e25ff848"
  identifier        = "campaignex-db"
  db_name           = "campaignex"
  engine            = "postgres"

  engine_version = "14.6"
  instance_class = "db.t3.micro"
  multi_az       = true

  username = "postgresAdmin"
  password = "pgDefaultPwdToChange!"

  skip_final_snapshot = true
}

resource "aws_db_instance" "profiler_db" {
  allocated_storage = 20
  storage_type      = "gp3"

  db_subnet_group_name = aws_db_subnet_group.database.name
  vpc_security_group_ids = [
    aws_security_group.angelowl_postgres.id
  ]

  storage_encrypted = true
  kms_key_id        = "arn:aws:kms:ap-southeast-1:276374573009:key/5e9b5264-8d3d-4496-8682-7268e25ff848"
  identifier        = "profiler-db"
  db_name           = "profiler"
  engine            = "postgres"

  engine_version = "14.6"
  instance_class = "db.t3.micro"
  multi_az       = true

  username = "postgresAdmin"
  password = "pgDefaultPwdToChange!"

  skip_final_snapshot = true
}
