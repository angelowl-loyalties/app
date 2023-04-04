data "aws_region" "current" {}

resource "aws_vpc_endpoint" "apigw_vpce" {
  private_dns_enabled = false
  security_group_ids = [
    aws_security_group.angelowl_http_s_ingress.id,
    aws_security_group.angelowl_outbound.id,
  ]

  service_name = "com.amazonaws.${data.aws_region.current.name}.execute-api"

  vpc_endpoint_type = "Interface"
  vpc_id            = aws_vpc.angelowl.id

  subnet_ids = [
    aws_subnet.angelowl_private_a.id,
    aws_subnet.angelowl_private_b.id,
    aws_subnet.angelowl_private_c.id
  ]

  tags = {
    "Name" = "api-gateway-vpce"
  }
}

resource "aws_api_gateway_vpc_link" "angelowl" {
  name        = "AngelOwl Ingress Link REST"
  description = "AngelOwl Ingress Link REST"
  target_arns = ["arn:aws:elasticloadbalancing:ap-southeast-1:276374573009:loadbalancer/net/angelowl-ingress/bb60e499bd17760e"]
}

resource "aws_api_gateway_rest_api" "angelowl_rest" {
  body = file("angelowl-rest-api.json")
  name              = "AngelOwl EKS APIGW REST"
  put_rest_api_mode = "overwrite"

  endpoint_configuration {
    types = ["REGIONAL"]
  }
}

data "aws_iam_policy_document" "restapi" {
  statement {
    effect = "Allow"

    principals {
      type        = "AWS"
      identifiers = ["*"]
    }

    actions   = ["execute-api:Invoke"]
    resources = [
        aws_api_gateway_rest_api.angelowl_rest.execution_arn,
        "${aws_api_gateway_rest_api.angelowl_rest.execution_arn}/*/*",
        "${aws_api_gateway_rest_api.angelowl_rest.execution_arn}/*/*/*",
    ]

    # condition {
    #   test     = "IpAddress"
    #   variable = "aws:SourceIp"
    #   values   = ["0.0.0.0/0"]
    # }
  }
}
# resource "aws_api_gateway_rest_api_policy" "test" {
#   rest_api_id = aws_api_gateway_rest_api.angelowl_rest.id
#   policy      = data.aws_iam_policy_document.restapi.json
# }

resource "aws_api_gateway_deployment" "default" {
  rest_api_id = aws_api_gateway_rest_api.angelowl_rest.id

  triggers = {
    redeployment = sha1(jsonencode(aws_api_gateway_rest_api.angelowl_rest.body))
  }

  lifecycle {
    create_before_destroy = true
  }
}


resource "aws_api_gateway_stage" "default" {
  deployment_id = aws_api_gateway_deployment.default.id
  rest_api_id   = aws_api_gateway_rest_api.angelowl_rest.id
  stage_name    = "default"
}

resource "aws_api_gateway_base_path_mapping" "base_path_mapping" {
  api_id      = aws_api_gateway_rest_api.angelowl_rest.id
  stage_name  = aws_api_gateway_stage.default.stage_name
  domain_name = aws_apigatewayv2_domain_name.itsag1t2.id
}

resource "aws_apigatewayv2_domain_name" "itsag1t2" {
  domain_name = "itsag1t2.com"

  domain_name_configuration {
    certificate_arn = "arn:aws:acm:ap-southeast-1:276374573009:certificate/b6d4ce5d-ed8f-46b5-88f4-3b9f8ca604b5"
    security_policy = "TLS_1_2"
    endpoint_type   = "REGIONAL"
  }
}
