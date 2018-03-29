# -*- coding: utf-8 -*-
# Generated by Django 1.11.3 on 2017-07-17 00:54
from __future__ import unicode_literals

from django.db import migrations, models
import django.db.models.deletion


class Migration(migrations.Migration):

    initial = True

    dependencies = [
        ('tipo_usuario', '0001_initial'),
    ]

    operations = [
        migrations.CreateModel(
            name='Trabajador',
            fields=[
                ('id', models.AutoField(primary_key=True, serialize=False)),
                ('nombre', models.CharField(max_length=255)),
                ('apellido', models.CharField(max_length=255)),
                ('username', models.CharField(max_length=20)),
                ('area', models.CharField(max_length=45)),
                ('dni', models.CharField(max_length=15, unique=True)),
                ('email', models.CharField(max_length=200)),
                ('sexo', models.CharField(max_length=40)),
                ('tipo_usuario', models.ForeignKey(on_delete=django.db.models.deletion.CASCADE, to='tipo_usuario.Tipo_Usuario')),
            ],
            options={
                'ordering': ['id'],
            },
        ),
    ]