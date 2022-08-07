import './style.css'
import * as PIXI from 'pixi.js'
import gowasm from './src/gowasm'


gowasm(() => {
    bunnyApp(PIXI)
})