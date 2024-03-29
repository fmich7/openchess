export default class GameController {
  constructor(
    public isRanked: boolean,
    public time: number,
    public timeAdded: number,
    public playerOneId: string,
    public playerTwoId: string,
    public gameType: string
  ) {
    this.whiteTime = time * 1000;
    this.blackTime = time * 1000;
    this.gameStatus = "Game in play!";
    this.isPlayerOneWhite = Math.random() < 0.5;
    this.playerOneToMove = this.isPlayerOneWhite;
  }

  whiteTime: number = 0;
  blackTime: number = 0;
  isPlayerOneWhite: boolean = false;
  playerOneToMove: boolean = false;
  gameStatus: string = "Game in play!";
}
