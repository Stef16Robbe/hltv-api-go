# HLTV-SCRAPING

## Random notes:
- use chromedp to start chrome engine,
- start it when hltv API object is created
- and I want to use defer cancel() so it can keep running and 
  we dont need to restart the chrome headless instance for every request we make
- we probably want to keep using the same context for all requests until we get rate limited
    - at that point we might need a new browser/tab/session, and perhaps another IP (but that's a problem for later :)
- getting blocked by cloudflare...
    - HUGE SHOUTOUT: github.com/Davincible/chromedp-undetected
- for getting complete body in chromedp Google gives us this issue;
    - https://github.com/chromedp/chromedp/issues/128#issuecomment-498051634
- BUT the following is a more stable solution:
    - https://github.com/chromedp/chromedp/issues/762#issuecomment-788836126
- TODO dl hltv html so we dont have to spam request their site so no ip block :)
- using: https://github.com/golang-standards/project-layout
- https://blog.boot.dev/golang/golang-enum/
- go install -v github.com/noho-digital/enumer
    - enumer -type=Map
    - enumer -type=Veto
    - etc...

## Ideas:
- recent map vetoes from both teams
  - want to know possible map picks
    - get map ban and pick %?
- team matchups, who has been beating the other more?
- player ratings, who stands out?
- who did these teams beat to get to this position in the tournament?
- an indication of teamwork?
  - flash assists? (naphony)
- introduce config: date range, team, matchType, lan yes/no
