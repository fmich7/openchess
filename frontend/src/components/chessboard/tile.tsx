const Tile = (props: { col: number; row: number; image?: string }) => {
  let tile_color = "";
  // const columns = "ABCDEFGHI";
  // const column = columns[props.col];

  // fills color of a tile
  if (Math.abs(props.col - props.row + 2) % 2 == 0) {
    tile_color = "bg-yellow-900";
  } else tile_color = "bg-orange-300";

  return (
    <div className={"grid place-content-center w-16 h-16 " + tile_color}>
      {props.image && (
        <div
          className="piece"
          style={{ backgroundImage: `url(${props.image})` }}
        ></div>
      )}
    </div>
  );
};

export default Tile;
