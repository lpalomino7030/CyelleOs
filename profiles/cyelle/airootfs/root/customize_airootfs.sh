#!/usr/bin/env bash

set -e

# Crear usuario live
useradd -G wheel,audio,video,storage,input -s /bin/bash cyelle

# Contraseña (solo para pruebas)
echo "cyelle:cyelle" | chpasswd

# Permitir sudo al grupo wheel
sed -i 's/^# %wheel ALL=(ALL:ALL) ALL/%wheel ALL=(ALL:ALL) ALL/' /etc/sudoers

# Crear carpetas de configuración
mkdir -p /home/cyelle/.config

chmod 644 /home/cyelle/.bash_profile

chown -R cyelle:cyelle /home/cyelle