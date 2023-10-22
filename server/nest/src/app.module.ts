import { Module } from '@nestjs/common';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { PrismaService } from './prisma.service';
import { MongooseModule } from '@nestjs/mongoose';
import { Board, BoardSchema } from './mongo/board.schema';
import { RedisService } from './redis.service';

@Module({
  imports: [
    MongooseModule.forRoot('mongodb://localhost/nest'),
    MongooseModule.forFeature([{ name: Board.name, schema: BoardSchema }]),
  ],
  controllers: [AppController],
  providers: [AppService, PrismaService, RedisService],
})
export class AppModule {}
