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

resource "aws_keyspaces_table" "transactions" {
  table_name    = "transactions"
  keyspace_name = aws_keyspaces_keyspace.angelowl.name

  comment {
    message = "Transactions table"
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
      name = "transaction_id"
      type = "ascii"
    }
  }
}

output "keyspaces_endpoint_dns" {
  value = aws_vpc_endpoint.keyspaces.dns_entry
}

# (
#     id               uuid primary key,
#     amount           double,
#     card_id          uuid,
#     card_pan         text,
#     card_type        text,
#     currency         text,
#     mcc              int,
#     merchant         text,
#     remarks          text,
#     reward_amount    double,
#     sgd_amount       double,
#     transaction_date date,
#     transaction_id   text
# );
