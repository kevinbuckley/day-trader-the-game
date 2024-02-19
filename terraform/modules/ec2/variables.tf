variable "subnet_id" {
    type  = string
    description = "Subnet to attach the network interface to for the ec2 instance"
}

variable "subnet_private_ips" { 
    type = list
    description = "private IPs for the network interface"
}

variable "ec2_ami_id" {
    type = string 
    description = "ami for the ec2 instance that is created"
}

variable "ec2_instance_type" {
    type = string
    description = "instance type for the ec2 isntance"
}
