'use client'
import sock from '../public/sock.png';
import Image from 'next/image';
import wsInstance from "@/app/net/web-socket";
import {ColorPicker} from "antd";
import {useState} from "react";

export default function Home() {
  const [color, setColor] = useState();
  const [canvasX, setCanvasX] = useState();
  const [canvasY, setCanvasY] = useState();

  const clickCanvas = (e) => {
    setCanvasX(e.clientX - e.currentTarget.offsetLeft);
    setCanvasY(e.clientY - e.currentTarget.offsetTop);
    wsInstance.send(JSON.stringify({CanvasX: canvasX, CanvasY: canvasY, Color: color}));
  }
  const selectColor = (nextColor) => {
    // debugger;
    const rgb = {
      R: Math.round(nextColor.metaColor.r),
      G: Math.round(nextColor.metaColor.g),
      B: Math.round(nextColor.metaColor.b)
    }
    setColor(rgb);
    // console.log(`nextColor: ${JSON.stringify(rgb)}`);
  }
  return (
    <main className="flex min-h-screen flex-col items-center">
      <div className={"text-xl font-bold flex flex-row"}>
        <div>pixel-socks</div>
        <div><Image className={"w-4 pt-1 ml-2"} src={sock} alt={"sock"} data-testid={"sock-image"}/></div>
      </div>
      <div className={"mt-24 flex flex-row"}>
        <div className={"mr-6 mt-9"} data-testid={"instruction-select-color"}>select a color -&gt;</div>
        <div className={"mr-6 mt-8"}>
          <ColorPicker onChange={selectColor} data-testid={"color-picker"}/>
        </div>
        <div>
          <canvas height={100} width={100} style={{background: 'white'}} onClick={clickCanvas} data-testid={"canvas"}></canvas>
        </div>
        {color && <div className={"ml-6 mt-9"} data-testid={"instruction-drop-pixel"}>&lt;- drop a pixel</div>}
      </div>
    </main>
  )
}
