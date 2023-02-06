# Create MSK instance
resource "aws_msk_cluster" "angelowl" {
  cluster_name           = "angelowl-msk"
  kafka_version          = "3.3.1"
  number_of_broker_nodes = 2

  encryption_info {
    encryption_in_transit {
      client_broker = "PLAINTEXT"
      in_cluster    = true
    }
  }

  broker_node_group_info {
    instance_type = "kafka.t3.small"

    client_subnets = [
      aws_subnet.angelowl_private_a.id,
      aws_subnet.angelowl_private_b.id
    ]

    security_groups = [
      aws_security_group.angelowl_kafka.id
    ]

    storage_info {
      ebs_storage_info {
        volume_size = 20
      }
    }
  }

  configuration_info {
    arn      = aws_msk_configuration.kafka_config.arn
    revision = aws_msk_configuration.kafka_config.latest_revision
  }

  enhanced_monitoring = "PER_BROKER"

  tags = {
    Name = "angelowl_msk"
  }
}

# Create MSK configuration
resource "aws_msk_configuration" "kafka_config" {
  name              = "angelowl-msk-config"
  kafka_versions    = ["3.3.1"]
  server_properties = <<PROPERTIES
log.retention.hours=168
auto.create.topics.enable=true
num.partitions=12
PROPERTIES
}

output "kafka_endpoint" {
  value = aws_msk_cluster.angelowl.bootstrap_brokers
}
