export const ShortToName = (short:string|string[]):string => {
    switch (short){
        case "lmop":
            return "Lost Mines of Phandelver"
        case "hotdq":
            return "Hoard of the Dragon Queen"
        case "bgdia":
            return "Baldurs Gate: Descent into Avernus"
        case "cm":
            return "Candlekeep Mysteries"
        case "pota":
            return "Princes of the Apocalypse"
        case "gos":
            return "Ghosts of the Saltmarsh"
        case "sacoc":
            return "Strixhaven: A Curriculum of Chaos"
        case "idrotf":
            return "Icewind Dale: Rime of the Frostmaiden"
        case "toa":
            return "Tomb of Annihilation"
        case "twbtw":
            return "The Wild before the Witchlight"
        case "wdh":
            return "Waterdeep: Dragon Heist"
        case "cos":
            return "Curse of Strahd"
        default:
            return "not found"
    }

}