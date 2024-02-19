

provider "random" {}

resource "random_pet" "instance" {
    length = 3
}

resource "aws_network_interface" "network_interface" {
    count       = length(var.subnet_private_ips)
    subnet_id   = var.subnet_id
    private_ips = [element(var.subnet_private_ips, count.index).ip]
}

resource "aws_instance" "ec2-instance" {
    count = length(var.subnet_private_ips)
    ami = var.ec2_ami_id 
    instance_type =  var.ec2_instance_type
    
    network_interface {
        network_interface_id = aws_network_interface.network_interface[count.index].id
        device_index         = 0
    }

    tags = { 
        Name = "${random_pet.instance.id}"
    }
}