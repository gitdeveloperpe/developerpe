# -*- coding: utf-8 -*-
# Generated by Django 1.11.4 on 2017-08-17 05:58
from __future__ import unicode_literals

from django.db import migrations, models
import django.db.models.deletion


class Migration(migrations.Migration):

    initial = True

    dependencies = [
        ('sala', '0004_auto_20170816_2346'),
    ]

    operations = [
        migrations.CreateModel(
            name='Entrada',
            fields=[
                ('id', models.AutoField(primary_key=True, serialize=False)),
                ('hora_compra', models.TimeField()),
                ('cantidad', models.IntegerField()),
                ('precio', models.DecimalField(decimal_places=4, default=0, max_digits=10)),
                ('total_pagado', models.DecimalField(decimal_places=4, default=0, max_digits=10)),
                ('pelicula', models.ForeignKey(on_delete=django.db.models.deletion.CASCADE, to='sala.Pelicula')),
                ('sala', models.ForeignKey(on_delete=django.db.models.deletion.CASCADE, to='sala.Sala')),
            ],
        ),
        migrations.CreateModel(
            name='Funcion',
            fields=[
                ('id', models.AutoField(primary_key=True, serialize=False)),
                ('hora_funcion', models.TimeField()),
                ('pelicula', models.ManyToManyField(to='sala.Pelicula')),
                ('sala', models.ForeignKey(on_delete=django.db.models.deletion.CASCADE, to='sala.Sala')),
            ],
        ),
        migrations.CreateModel(
            name='TipoEntrada',
            fields=[
                ('id', models.AutoField(primary_key=True, serialize=False)),
                ('tipo', models.CharField(max_length=255)),
            ],
        ),
        migrations.AddField(
            model_name='entrada',
            name='tipo_entrada',
            field=models.ManyToManyField(to='compras.TipoEntrada'),
        ),
    ]
