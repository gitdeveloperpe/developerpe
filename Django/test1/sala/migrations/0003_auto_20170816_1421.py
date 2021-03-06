# -*- coding: utf-8 -*-
# Generated by Django 1.11.4 on 2017-08-16 19:21
from __future__ import unicode_literals

from django.db import migrations, models


class Migration(migrations.Migration):

    initial = True

    dependencies = [
        ('sala', '0002_delete_persona'),
    ]

    operations = [
        migrations.CreateModel(
            name='Butaca',
            fields=[
                ('id', models.AutoField(primary_key=True, serialize=False)),
                ('numero', models.CharField(max_length=10)),
            ],
        ),
        migrations.CreateModel(
            name='Categoria',
            fields=[
                ('id', models.AutoField(primary_key=True, serialize=False)),
                ('categoria', models.CharField(max_length=255)),
            ],
        ),
        migrations.CreateModel(
            name='Fila',
            fields=[
                ('id', models.AutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('fila', models.CharField(max_length=1, unique=True)),
            ],
        ),
        migrations.CreateModel(
            name='Pelicula',
            fields=[
                ('id', models.AutoField(primary_key=True, serialize=False)),
                ('nombre', models.CharField(max_length=255)),
                ('nombre_alterno', models.CharField(blank=True, max_length=255)),
                ('duracion', models.CharField(max_length=40)),
                ('fecha_estreno', models.DateField()),
                ('idioma_original', models.CharField(max_length=255)),
                ('director', models.CharField(max_length=255)),
                ('sinopsis', models.TextField()),
                ('categoria', models.ManyToManyField(to='sala.Categoria')),
            ],
        ),
        migrations.CreateModel(
            name='Sala',
            fields=[
                ('id', models.AutoField(primary_key=True, serialize=False)),
                ('numero', models.CharField(max_length=2)),
                ('tipo_sala', models.CharField(max_length=2)),
                ('butacas', models.ManyToManyField(to='sala.Butaca')),
                ('pelicula', models.ManyToManyField(to='sala.Pelicula')),
            ],
        ),
        migrations.AddField(
            model_name='butaca',
            name='fila',
            field=models.ManyToManyField(to='sala.Fila'),
        ),
    ]
