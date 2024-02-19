variable "region" {
  type        = string
  description = "AWS region to use"
  default     = "us-east-1"
}
variable "project_name" {
  type        = string
  description = "name of the project"
  default     = "kbux-infrastructure"
}
variable "ec2_ips" {
  type        = list(object({ ip = string, type = string }))
  description = "list of ec2 instance ips"
  default = [
    { type = "t2.micro", ip = "172.16.10.105", create_network = true },
    { type = "t2.micro", ip = "172.16.10.104", create_network = true }
  ]
}