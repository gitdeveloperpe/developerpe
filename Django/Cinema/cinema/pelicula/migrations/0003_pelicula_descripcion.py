# -*- coding: utf-8 -*-
# Generated by Django 1.11.3 on 2017-08-01 04:42
from __future__ import unicode_literals

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('pelicula', '0002_auto_20170731_2334'),
    ]

    operations = [
        migrations.AddField(
            model_name='pelicula',
            name='descripcion',
            field=models.TextField(default=1),
            preserve_default=False,
        ),
    ]