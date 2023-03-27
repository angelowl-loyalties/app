resource "aws_route53_zone" "primary" {
  name = "itsag1t2.com"
}

resource "aws_route53_record" "primary" {
  zone_id = aws_route53_zone.primary.zone_id
  name    = "itsag1t2.com"
  type    = "A"

  alias {
    name                   = aws_apigatewayv2_domain_name.itsag1t2.domain_name_configuration[0].target_domain_name
    zone_id                = aws_apigatewayv2_domain_name.itsag1t2.domain_name_configuration[0].hosted_zone_id
    evaluate_target_health = false
  }
}

resource "aws_route53_record" "helm_dashboard" {
  zone_id = aws_route53_zone.primary.zone_id
  name    = "helm.dashboard.itsag1t2.com"
  type    = "CNAME"
  ttl     = "30"
  records = [var.nginx_ingress_lb_ip]
}

resource "aws_route53_record" "kube_dashboard" {
  zone_id = aws_route53_zone.primary.zone_id
  name    = "kube.dashboard.itsag1t2.com"
  type    = "CNAME"
  ttl     = "30"
  records = [var.nginx_ingress_lb_ip]
}

resource "aws_route53_record" "kafka_dashboard" {
  zone_id = aws_route53_zone.primary.zone_id
  name    = "kafka.dashboard.itsag1t2.com"
  type    = "CNAME"
  ttl     = "30"
  records = [var.nginx_ingress_lb_ip]
}

resource "aws_route53_record" "campaignex" {
  zone_id = aws_route53_zone.primary.zone_id
  name    = "campaignex.itsag1t2.com"
  type    = "CNAME"
  ttl     = "60"
  records = [var.nginx_ingress_lb_ip]
}

resource "aws_route53_record" "profiler" {
  zone_id = aws_route53_zone.primary.zone_id
  name    = "profiler.itsag1t2.com"
  type    = "CNAME"
  ttl     = "60"
  records = [var.nginx_ingress_lb_ip]
}

resource "aws_route53_record" "informer" {
  zone_id = aws_route53_zone.primary.zone_id
  name    = "informer.itsag1t2.com"
  type    = "CNAME"
  ttl     = "60"
  records = [var.nginx_ingress_lb_ip]
}
