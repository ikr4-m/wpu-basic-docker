import type { Response } from 'express';
import { Controller, Get, Res, StreamableFile } from '@nestjs/common';
import { join } from 'path';
import { createReadStream } from 'fs';

@Controller()
export class AppController {
  @Get()
  getWarmVideo(@Res({ passthrough: true }) res: Response): StreamableFile {
    res.set({
      'Content-Type': 'video/mp4',
      'Content-Disposition': 'inline; filename="video.mp4"',
    });
    const stream = createReadStream(join(process.cwd(), 'files', 'video.mp4'))
    return new StreamableFile(stream);
  }
}
