locals {
  ztest_local_map = {
    "test" = "1234"
  }

  test_list = ["foo", "bar", "doo"]
  name      = "name"

  map = {
    name = local.name
    list = local.test_list
  }
}

module "test" {
  name   = local.name
  count  = 1
  source = "./test"
}

data "aws_eks_cluster" "test" {
  name = local.name
}

provider "aws" {
  region = "eu-west-1"
}

module "btest" {
  // source = "git@github.com:CCV-Group/terraform-aws-ccv-eks.git?ref=sandbox"
  source = "./test-b"
  name   = local.mtest.name
  zname  = "zname"
  create = local.mtest.create_eks
  user   = local.cluster_enabled_log_types
  list   = module.martijn_sandbox_vpc[0].private_subnets_cidr_blocks
  defaults = {
    name = ""
    avar = false
  }
}
module "eks" {
  example = {
    create_launch_template = true
    desired_capacity       = 1
    max_capacity           = 10
    min_capacity           = 1

    disk_size       = 50
    disk_type       = "gp3"
    disk_throughput = 150
    disk_iops       = 3000

    instance_types = ["t3.large"]
    capacity_type  = "SPOT"

    bootstrap_env = {
      CONTAINER_RUNTIME = "containerd"
      USE_MAX_PODS      = false
    }
    kubelet_extra_args = "--max-pods=110"
    k8s_labels = {
      GithubRepo = "terraform-aws-eks"
      GithubOrg  = "terraform-aws-modules"
    }
    additional_tags = {
      ExtraTag = "example2"
    }
    taints = [
      {
        key    = "dedicated"
        value  = "gpuGroup"
        effect = "NO_SCHEDULE"
      }
    ]
    update_config = {
      max_unavailable_percentage = 50 # or set `max_unavailable`
    }
  }
}

