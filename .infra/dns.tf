resource "aws_route53_zone" "primary" {
  name = "itsag1t2.com"
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
