import { Injectable, OnModuleInit } from '@nestjs/common';
import Redis from 'ioredis';

@Injectable()
export class RedisService implements OnModuleInit {
  private client: Redis;

  onModuleInit() {
    this.client = new Redis({
      host: 'devcontainer',
      port: 6379,
    });
  }

  onModuleDestroy() {
    this.client.disconnect();
  }

  async set(key: string, value: any, ttl: number) {
    await this.client.set(key, JSON.stringify(value), 'EX', ttl);
  }

  async get(key: string) {
    const data = await this.client.get(key);
    return data ? JSON.parse(data) : null;
  }

  async delete(key: string) {
    await this.client.del(key);
  }
}
