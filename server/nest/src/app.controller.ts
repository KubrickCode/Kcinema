import { Body, Controller, Get, Post } from '@nestjs/common';
import { AppService } from './app.service';

@Controller()
export class AppController {
  constructor(private readonly appService: AppService) {}

  @Get()
  getHello(): string {
    return this.appService.getHello();
  }

  @Get('users')
  async getUsers() {
    return await this.appService.getUsers();
  }

  @Post('user')
  async createUser(@Body() user: { email: string }) {
    return await this.appService.createUser(user.email);
  }

  @Get('boards')
  async getBoards() {
    return await this.appService.getBoards();
  }

  @Post('board')
  async createBoard(@Body() board: { title: string }) {
    return await this.appService.createBoard(board.title);
  }

  @Get('messages')
  async getMessage() {
    return await this.appService.getMessage();
  }

  @Post('message')
  async createMessage(@Body() message: { message: string }) {
    return await this.appService.createMessage(message.message);
  }
}
