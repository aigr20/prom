function padZero(value: number): string {
  return value <= 9 ? `0${value}` : String(value);
}

export function formatDate(date: Date): string {
  date = new Date(date);
  const year = date.getFullYear();
  const month = date.getMonth() + 1;
  const day = date.getDate() + 1;
  const hour = date.getHours();
  const minute = date.getMinutes();

  return `${year}-${month}-${day} ${padZero(hour)}:${padZero(minute)}`;
}
