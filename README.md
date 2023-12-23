# Go Speedtest

A network speed testing application built on top of Ookla's SpeedTest CLI tooling. Also supports weather gathering(for satellite Wi-Fi users).

## Usage

1) Ensure that the [ookla speedtest cli](https://www.speedtest.net/apps/cli) is installed
2) You can just run it this tool as a simple binary: `./go_speedtest`
3) By default, to determine weather, it will use the location given by your ip-address.
However, this might not be useful if your ip-address does not map to your correct location.
You can instead pass in your latitude, longitude coordinates: `./go_speedtest <latitude> <longitude>`.

## Motivation

Satellite internet is subject to environmental factors like weather.
I'm building this tool to help better understand the consequences of things like weather on network speed.

## Weather
Pulls from https://www.weatherapi.com/.

## Backend
Uses a simple SQLite backend(at least for now).