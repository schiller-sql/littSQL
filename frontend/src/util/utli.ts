const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ";

function basicLetterFromNumber(n) {
  return alphabet[n];
}

export function letterFromNumber(n) {
  if (n === 0) return "A";
  let name = "";
  while (n >= 0) {
    if (n >= alphabet.length) {
      n -= alphabet.length - 1;
      name += "Z";
    } else if (n !== 0) {
      name += basicLetterFromNumber(n);
      break;
    }
  }
  return name;
}
