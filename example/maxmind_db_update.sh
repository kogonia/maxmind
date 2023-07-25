#! /usr/bin/env bash

license_key="YOUR_KEY_HERE"

archive_name="GeoLite2-ASN-CSV"
db_file="GeoLite2-ASN-Blocks-IPv4.csv"

wget "https://download.maxmind.com/app/geoip_download?edition_id=${archive_name}&license_key=${license_key}&suffix=zip" -O "${archive_name}.zip"
unzip -p "${archive_name}.zip" "*/${db_file}" > "${db_file}"
rm -f "${archive_name}.zip"
