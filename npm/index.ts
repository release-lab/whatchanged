import ffi from "ffi-napi";

export interface Options {
  version: string[];
  branch: string;
  format: string;
  skip_format?: boolean;
  preset?: string;
  template_file?: string;
  template?: string;
  silent?: boolean;
}

export class WhatchangedFFI {
  #lib: any;

  /**
   *
   * @description Linux: whatchanged.so
   * @description Windows: whatchanged.dll
   * @description macOS: whatchanged.dylib
   * @param libPath whatchanged Dynamic-link library filepath without extension
   */
  constructor(private libPath: string) {
    this.#lib = ffi.Library(libPath, {
      Generate: ["string", ["string", "string"]],
    });
  }

  public generate(repo: string, options: Options): Promise<string> {
    return new Promise((resolve, reject) => {
      this.#lib.Generate.async(
        repo,
        JSON.stringify(options),
        (err: Error, res: string | null) => {
          if (err) {
            return reject(err);
          }

          if (res === null) return reject(new Error(""));

          const { result: changelog, error } = JSON.parse(res);
          if (error) {
            return reject(new Error(error));
          }

          resolve(changelog);
        }
      );
    });
  }
}
