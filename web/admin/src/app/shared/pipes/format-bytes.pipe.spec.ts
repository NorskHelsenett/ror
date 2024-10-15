import { FormatBytesPipe } from './format-bytes.pipe';

describe('FormatBytesPipe', () => {
  it('create an instance', () => {
    const pipe = new FormatBytesPipe(null);
    expect(pipe).toBeTruthy();
  });
});
