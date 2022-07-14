FROM redis
COPY redis.conf /usr/local/etc/redis/redis.conf
CMD [ "redis-server", "/usr/local/etc/redis/redis.conf" ]
EXPOSE 6380
# docker login
# docker build -t edumar111/fastpv-redis:v1.0.0 -f Dockerfile .
# docker push     edumar111/fastpv-redis:v1.0.0

# docker run --name fastpv-redis -d -p 6380:6380  edumar111/fastpv-redis:v1.0.0