export default class GameController {
  playerOneId: string; // players id
  playerTwoId: string;
  time: number; // time
  timeAdded: number;
  whiteTime: number; // clock time
  blackTime: number;
  isPlayerOneWhite: boolean; // is player one starting with white
  playerOneToMove: boolean;
  // misc
  gameStatus: string; // Started finished etc,
  isRanked: boolean; // affects elo
  gameType: string; // bullet etc.

  constructor(
    isRanked: boolean,
    time: number,
    timeAdded: number,
    playerOneId: string,
    playerTwoId: string,
    gameType: string
  ) {
    this.time = time;
    this.timeAdded = timeAdded;
    this.whiteTime = time;
    this.blackTime = time;
    this.playerOneId = playerOneId;
    this.playerTwoId = playerTwoId;
    this.gameStatus = "Game in play!";
    this.isRanked = isRanked;
    this.gameType = gameType;
    // assign players to sides
    this.isPlayerOneWhite = Math.random() < 0.5 ? true : false;
    this.playerOneToMove = this.isPlayerOneWhite;
  }
}
