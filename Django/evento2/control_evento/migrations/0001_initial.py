# -*- coding: utf-8 -*-
# Generated by Django 1.10 on 2017-09-09 16:15
from __future__ import unicode_literals

from django.db import migrations, models


class Migration(migrations.Migration):

    initial = True

    dependencies = [
        ('staff', '0001_initial'),
    ]

    operations = [
        migrations.CreateModel(
            name='Asistente',
            fields=[
                ('id', models.AutoField(primary_key=True, serialize=False)),
                ('dni', models.CharField(max_length=10)),
                ('nombre', models.CharField(max_length=70)),
                ('apellidos', models.CharField(max_length=200)),
                ('email', models.EmailField(max_length=254)),
                ('direccion', models.TextField()),
                ('telefono', models.CharField(max_length=9)),
            ],
        ),
        migrations.CreateModel(
            name='Evento',
            fields=[
                ('id', models.AutoField(primary_key=True, serialize=False)),
                ('fecha_inicio', models.DateField()),
                ('fecha_fin', models.DateField()),
                ('asistentes', models.ManyToManyField(to='control_evento.Asistente')),
            ],
        ),
        migrations.CreateModel(
            name='Seminario',
            fields=[
                ('id', models.AutoField(primary_key=True, serialize=False)),
                ('tema', models.CharField(max_length=80)),
                ('fecha', models.DateField()),
                ('fecha_inicio', models.DateTimeField()),
                ('fecha_fin', models.DateTimeField()),
                ('ponente', models.ManyToManyField(to='staff.Ponente')),
            ],
        ),
        migrations.AddField(
            model_name='evento',
            name='seminario',
            field=models.ManyToManyField(to='control_evento.Seminario'),
        ),
    ]
