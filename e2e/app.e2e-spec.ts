import { Angular2JwtStarterPage } from './app.po';

describe('angular2-jwt-starter App', function() {
  let page: Angular2JwtStarterPage;

  beforeEach(() => {
    page = new Angular2JwtStarterPage();
  });

  it('should display message saying app works', () => {
    page.navigateTo();
    expect(page.getParagraphText()).toEqual('app works!');
  });
});
