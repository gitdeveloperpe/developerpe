# -*- coding: utf-8 -*-
# Generated by Django 1.11.4 on 2017-08-18 19:32
from __future__ import unicode_literals

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('compras', '0001_initial'),
    ]

    operations = [
        migrations.AlterField(
            model_name='entrada',
            name='precio',
            field=models.DecimalField(decimal_places=4, default=7, max_digits=10),
        ),
    ]