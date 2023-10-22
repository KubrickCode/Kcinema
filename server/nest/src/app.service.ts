import { Injectable } from '@nestjs/common';
import { PrismaService } from './prisma.service';
import { InjectModel } from '@nestjs/mongoose';
import { Board } from './mongo/board.schema';
import { Model } from 'mongoose';
import { RedisService } from './redis.service';

@Injectable()
export class AppService {
  constructor(
    private prisma: PrismaService,
    @InjectModel(Board.name)
    private boardModel: Model<Board>,
    private redis: RedisService,
  ) {}
  getHello(): string {
    return 'Hello World!';
  }

  async createUser(email: string) {
    await this.prisma.user.create({ data: { email } });
  }

  async getUsers() {
    return await this.prisma.user.findMany({});
  }

  async createBoard(title: string) {
    const createdBoard = new this.boardModel({ title });
    return createdBoard.save();
  }

  async getBoards() {
    return this.boardModel.find().exec();
  }

  async createMessage(message: string) {
    await this.redis.set('message', message, 10000);
  }

  async getMessage() {
    return await this.redis.get('message');
  }
}
