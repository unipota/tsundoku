export enum SNS {
  Twitter = 'twitter',
  Line = 'line',
  Facebook = 'facebook'
}

const hashTag = 'ツンドク'
const twitterId = 'tsundokuApp'

export const getShareText = (
  tsundokuPrice: number,
  booksCount: number
): string => {
  return `${booksCount}冊、${tsundokuPrice.toLocaleString(
    'ja-JP'
  )}円分の積読をしています！`
}

export const getTargetURL = (
  snsName: SNS,
  shareText: string,
  shareURL: string
): string => {
  let targetURL = ''

  switch (snsName) {
    case SNS.Twitter:
      targetURL = `https://twitter.com/intent/tweet?url=${encodeURI(
        shareURL
      )}&text=${encodeURI(shareText)}&via=${twitterId}&hashtags=${encodeURI(
        hashTag
      )}`
      break
    case SNS.Line:
      targetURL = `https://social-plugins.line.me/lineit/share?url=${encodeURI(
        shareURL
      )}`
      break
    case SNS.Facebook:
      targetURL = `https://www.facebook.com/sharer/sharer.php?u=${encodeURI(
        shareURL
      )}`
      break
  }
  return targetURL
}
