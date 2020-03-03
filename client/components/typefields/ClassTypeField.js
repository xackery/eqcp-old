const ClassTypeField = ({value}) => {
  let classes = []
  if (value & (1 << 0) > 0) classes.push("WAR")
  if (value & (1 << 1) > 0) classes.push("CLR")
  if (value & (1 << 2) > 0) classes.push("PAL")
  if (value & (1 << 3) > 0) classes.push("RNG")
  if (value & (1 << 4) > 0) classes.push("SHD")
  if (value & (1 << 5) > 0) classes.push("DRU")
  if (value & (1 << 6) > 0) classes.push("MNK")
  if (value & (1 << 7) > 0) classes.push("BRD")
  if (value & (1 << 8) > 0) classes.push("ROG")
  if (value & (1 << 8) > 0) classes.push("ROG")
  if (value & (1 << 9) > 0) classes.push("SHM")
  if (value & (1 << 10) > 0) classes.push("NEC")
  if (value & (1 << 11) > 0) classes.push("WIZ")
  if (value & (1 << 12) > 0) classes.push("MAG")
  if (value & (1 << 13) > 0) classes.push("ENC")
  if (value & (1 << 14) > 0) classes.push("BST")
  if (value & (1 << 15) > 0) classes.push("BER")
  return classes.length === 17 ? "ALL" : classes.join(" ")
}

export default ClassTypeField