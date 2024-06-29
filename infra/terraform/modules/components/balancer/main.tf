resource "yandex_compute_instance" "balancer" {
  name                = var.name
  resources {
    cores             = var.cores
    memory            = var.memory
    }
  boot_disk {
    initialize_params {
      size            = var.size
      image_id        = "fd88bokmvjups3o0uqes"
      }
    }
  network_interface {
    subnet_id         = var.subnet_id
    }
  }
