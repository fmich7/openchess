type ChessGame = {
  ID: number;
  HostID: number;
  WhitePlayerID: number;
  BlackPlayerID: number;
  WhiteToMove: boolean;
  GameFEN: string;
  GameType: string;
  GameStatus: string;
  GameEnded: boolean;
  GameWonByWhite: boolean;
  IsRanked: boolean;
  Time: number;
  TimeAdded: number;
  MovesCount: number;
  MoveHistory: string;
  CreatedAt: string;
};

export type { ChessGame };
