resource "aws_vpc" "kbux_infra_vpc" {
  cidr_block = "172.16.0.0/16"
  tags = {
    Name = "kbux-infra"
  }
}


resource "aws_subnet" "kbux_infra_subnet" {
  vpc_id            = aws_vpc.kbux_infra_vpc.id
  cidr_block        = "172.16.10.0/24"
  availability_zone = "us-east-1d"
  tags = {
    Name = "kbux-infra"
  }
}
 
module "ec2-instance" {
  source = "./modules/ec2"

  subnet_id          = aws_subnet.kbux_infra_subnet.id
  subnet_private_ips = var.ec2_ips
  ec2_ami_id         = "ami-0c7217cdde317cfec"
  ec2_instance_type  = "t2.micro"
}