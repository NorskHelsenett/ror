import { HighlightDifferencePipe } from './highlight-difference.pipe';

describe('HighlightDifferencePipe', () => {
  it('create an instance', () => {
    const pipe = new HighlightDifferencePipe(null);
    expect(pipe).toBeTruthy();
  });
});
