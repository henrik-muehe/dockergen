Dockergen
=========

Allows generating an nginx config from `docker ps`. I run this periodically to expose services. Per default, every docker container exposing its private port 8080 is exposed. Modify the inline template and run using cron, for instance:

	*/5 * * * * go run /path/to/dockergen.go > /etc/nginx/sites-enabled/docker && service nginx reload


LICENSE
=======
This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
