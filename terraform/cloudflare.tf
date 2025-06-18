module "cloudflare_record" {
  source                     = "./modules/cloudflare"
  cloudflare_zone_id         = var.cloudflare_zone_id
  cloudflare_sub_domain_name = var.cloudflare_sub_domain_name
  cloudflare_record_content  = module.api_custom_domain.api_regional_domain_name
}
