# -*- coding: utf-8 -*-

from setuptools import setup

version = '1.0.0'

setup(
	name = 'pydemo',
	version = version,
	packages = ['pydemo'],
	entry_points = {
		'console_scripts': [
			'pydemo = pydemo:Run',
			'pydemo2 = pydemo:Run2'
		]
	},
	install_requires = [],
	description = "pydemo: a python package demo",
	author = 'tongsq',
	classifiers = [],
)
