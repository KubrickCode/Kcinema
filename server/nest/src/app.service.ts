import { Injectable } from '@nestjs/common';
import { PrismaService } from './prisma.service';

@Injectable()
export class AppService {
  constructor(private prisma: PrismaService) {}
  getHello(): string {
    return 'Hello World!';
  }

  async createUser(email: string) {
    await this.prisma.user.create({ data: { email } });
  }

  async getUsers() {
    return await this.prisma.user.findMany({});
  }
}
