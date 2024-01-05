export default class GameController {
  whiteTime: number;
  blackTime: number;

  constructor(time: number) {
    this.whiteTime = time;
    this.blackTime = time;
  }
}
