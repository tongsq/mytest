#!/usr/bin/env python3
# -*- coding: utf-8 -*-

Dict = {0: 'stop', 201: 'restart', 202: 'fresh-restart'}
print(Dict.get(201, 'system-exit'), Dict.get(100, 'system-exit'))
