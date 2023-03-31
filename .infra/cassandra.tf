resource "aws_keyspaces_keyspace" "angelowl" {
  name = "angelowl"
}

resource "aws_vpc_endpoint" "keyspaces" {
  vpc_id       = aws_vpc.angelowl.id
  service_name = "com.amazonaws.ap-southeast-1.cassandra"

  vpc_endpoint_type = "Interface"
  security_group_ids = [
    aws_security_group.angelowl_keyspaces.id
  ]

  subnet_ids = [
    aws_subnet.angelowl_private_a.id,
    aws_subnet.angelowl_private_b.id
  ]

  tags = {
    Name = "angelowl-keyspaces"
  }
}

resource "aws_keyspaces_table" "rewards" {
  table_name    = "rewards"
  keyspace_name = aws_keyspaces_keyspace.angelowl.name

  comment {
    message = "Table for rewards issued."
  }

  capacity_specification {
    throughput_mode = "PAY_PER_REQUEST"
  }

  schema_definition {
    partition_key {
      name = "card_id"
    }

    clustering_key {
      name     = "transaction_date"
      order_by = "DESC"
    }

    clustering_key {
      name = "id"
      order_by = "ASC"
    }

    static_column {
      name = "card_pan"
    }

    static_column {
      name = "card_type"
    }

    column {
      name = "id"
      type = "uuid"
    }

    column {
      name = "amount"
      type = "double"
    }

    column {
      name = "card_id"
      type = "uuid"
    }

    column {
      name = "card_pan"
      type = "ascii"
    }

    column {
      name = "card_type"
      type = "ascii"
    }

    column {
      name = "currency"
      type = "ascii"
    }

    column {
      name = "mcc"
      type = "int"
    }

    column {
      name = "merchant"
      type = "ascii"
    }

    column {
      name = "remarks"
      type = "text"
    }

    column {
      name = "reward_amount"
      type = "double"
    }

    column {
      name = "sgd_amount"
      type = "double"
    }

    column {
      name = "transaction_date"
      type = "date"
    }

    column {
      name = "created_at"
      type = "date"
    }

    column {
      name = "transaction_id"
      type = "ascii"
    }
  }
}

# CREATE TABLE angelowl.rewards (
#     card_id uuid,
#     transaction_date date,
#     amount double,
#     currency ascii,
#     id uuid,
#     mcc int,
#     merchant ascii,
#     remarks text,
#     reward_amount double,
#     sgd_amount double,
#     transaction_id ascii,
#     card_pan ascii static,
#     card_type ascii static,
#     PRIMARY KEY (card_id, transaction_date)
# ) WITH CLUSTERING ORDER BY (transaction_date DESC)

output "keyspaces_endpoint_dns" {
  value = aws_vpc_endpoint.keyspaces.dns_entry
}

