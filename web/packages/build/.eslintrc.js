/*
 * Teleport
 * Copyright (C) 2023  Gravitational, Inc.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

const path = require('path');
const tsConfigBase = require(
  path.join(path.resolve(__dirname, '../../..'), 'tsconfig.base.json')
);

const appPackages = ['teleport', 'e-teleport', 'teleterm'];
const libraryPackages = Object.keys(tsConfigBase.compilerOptions.paths)
  .map(p => p.split('/')[0])
  .filter(p => !appPackages.includes(p) && !p.startsWith('gen-proto-'));

const excludedPackagesPattern = [...appPackages, ...libraryPackages].join('|');

// importSortGroups are passed to `simple-import-sort` plugin to sort imports. Each group is an array of regexes.
// Imports are sorted by the first group that matches them, with newlines inserted between groups.
const importSortGroups = [
  // Side-effect imports (e.g. CSS)
  ['^\u0000(?!.*\u0000$)'],
  // Any 3rd-party imports (e.g. 'react', 'styled-components')
  [`^(?![./])(?!(?:${excludedPackagesPattern})(?:/|$))(?!.*\u0000$).+`],
  // Our library packages (e.g. 'shared', 'design')
  [`^(?:${libraryPackages.join('|')})(?:/)?(?!.*\u0000$)`],
  // Our app packages (e.g. '(e-)?teleport', 'teleterm')
  [`^(?:${appPackages.join('|')})(?:/)?(?!.*\u0000$)`],
  // Imports from 'gen-proto-ts/js'
  ['^gen-proto-(?:j|t)s(?:/)?(?!.*\u0000$)'],
  // Absolute and relative imports
  ['^[^.](?!.*\u0000$)', '^\\.(?!.*\u0000$)'],
  // Type imports
  ['\u0000$'],
];

module.exports = {
  parser: '@typescript-eslint/parser',
  parserOptions: {
    ecmaVersion: 6,
    ecmaFeatures: {
      jsx: true,
      experimentalObjectRestSpread: true,
    },
  },
  env: {
    es6: true,
    browser: true,
    node: true,
  },
  globals: {
    expect: true,
    jest: true,
  },
  ignorePatterns: ['**/dist/**'],
  extends: [
    'eslint:recommended',
    'plugin:@typescript-eslint/recommended',
    'plugin:import/errors',
    'plugin:import/warnings',
    'plugin:import/typescript',
  ],
  plugins: ['react', 'babel', 'import', 'simple-import-sort', 'react-hooks'],
  overrides: [
    {
      files: ['**/*.test.{ts,tsx,js,jsx}'],
      plugins: ['jest'],
      extends: [
        'plugin:jest/recommended',
        'plugin:testing-library/react',
        'plugin:jest-dom/recommended',
      ],
      rules: {
        'jest/prefer-called-with': 0,
        'jest/prefer-expect-assertions': 0,
        'jest/consistent-test-it': 0,
        'jest/no-try-expect': 0,
        'jest/no-hooks': 0,
        'jest/no-disabled-tests': 0,
        'jest/prefer-strict-equal': 0,
        'jest/prefer-inline-snapshots': 0,
        'jest/require-top-level-describe': 0,
        'jest/no-large-snapshots': [1, { maxSize: 200 }],
      },
    },
    // Allow `require` imports in .js files, as migrating our project to ESM modules requires a lot of
    // changes.
    {
      files: ['**/*.js'],
      rules: {
        '@typescript-eslint/no-require-imports': 1,
      },
    },
    // Allow `require` imports & no import sorting in this file
    {
      files: ['**/.eslintrc.js'],
      rules: {
        '@typescript-eslint/no-require-imports': 0,
        'simple-import-sort/imports': 0,
        'import/newline-after-import': 0,
      },
    },
  ],
  // Severity should be one of the following:
  // "off" or 0 - turn the rule off
  // "warn" or 1 - turn the rule on as a warning (doesn’t affect exit code)
  // "error" or 2 - turn the rule on as an error (exit code is 1 when triggered)
  rules: {
    // Sort imports/exports
    'import/order': 0,
    'simple-import-sort/imports': [2, { groups: importSortGroups }],
    'simple-import-sort/exports': 1,
    // Make sure imports are first
    'import/first': 2,
    // Add newline after all imports
    'import/newline-after-import': 2,
    // Merge duplicate imports (duplicate type-only imports are allowed)
    'import/no-duplicates': 2,

    // typescript-eslint recommends to turn import/no-unresolved off.
    // https://typescript-eslint.io/troubleshooting/typed-linting/performance/#eslint-plugin-import
    'import/no-unresolved': 0,
    'no-unused-vars': 0, // disabled to allow the typescript one to take over and avoid errors in reporting
    '@typescript-eslint/no-unused-vars': [
      2,
      // Allow unused vars/args that start with an underscore
      {
        varsIgnorePattern: '^_',
        argsIgnorePattern: '^_',
      },
    ],
    'no-unused-expressions': 0,
    '@typescript-eslint/no-unused-expressions': [
      2,
      { allowShortCircuit: true, allowTernary: true, enforceForJSX: true },
    ],
    '@typescript-eslint/no-empty-object-type': [
      2,
      // with-single-extends is needed to allow for interface extends like we have in jest.d.ts.
      { allowInterfaces: 'with-single-extends' },
    ],

    // <TODO> Enable these rules after fixing all existing issues
    '@typescript-eslint/no-use-before-define': 0,
    '@typescript-eslint/indent': 0,
    '@typescript/no-use-before-define': 0,
    '@typescript-eslint/no-var-requires': 0,
    '@typescript-eslint/camelcase': 0,
    '@typescript-eslint/class-name-casing': 0,
    '@typescript-eslint/no-explicit-any': 0,
    '@typescript-eslint/explicit-function-return-type': 0,
    '@typescript-eslint/explicit-member-accessibility': 0,
    '@typescript-eslint/prefer-interface': 0,
    '@typescript-eslint/no-empty-function': 0,
    '@typescript-eslint/no-this-alias': 0,
    '@typescript-eslint/no-unnecessary-type-assertion': 0,
    // Enable once existing issues are fixed (require 'project' in parserOptions)
    // - these catch a lot of easy-to-overlook issues.
    '@typescript-eslint/no-implied-eval': 0, // don't allow implied eval (e.g. passing strings to setTimeout)
    '@typescript-eslint/no-unsafe-unary-minus': 0,
    '@typescript-eslint/require-await': 0, // require async functions to have an await statement
    '@typescript-eslint/no-for-in-array': 0,
    '@typescript-eslint/await-thenable': 0, // don't allow awaiting non-promises
    '@typescript-eslint/no-array-delete': 0, // don't allow deleting array elements using `delete`
    // Likewise, but more cases to fix first.
    '@typescript-eslint/return-await': 0, // require awaiting promises in current scope before returning
    '@typescript-eslint/no-floating-promises': 0, // require await or return for promises
    '@typescript-eslint/no-misused-promises': 0,

    // </TODO>
    'comma-dangle': 0,
    'no-mixed-spaces-and-tabs': 0,
    'no-alert': 0,
    'import/no-named-as-default': 0,
    'import/default': 2,
    'no-underscore-dangle': 0,
    'no-case-declarations': 0,
    'prefer-const': 0,
    'no-var': 0,
    'prefer-rest-params': 0,
    'prefer-spread': 0,

    strict: 0,
    'no-console': 1,
    'no-trailing-spaces': 2,
    'react/display-name': 0,
    'react/jsx-no-undef': 2,
    'react/jsx-pascal-case': 2,
    'react/no-danger': 2,
    'react/jsx-no-duplicate-props': 2,
    'react/jsx-sort-prop-types': 0,
    'react/jsx-sort-props': 0,
    'react/jsx-uses-vars': 1,
    'react/no-did-mount-set-state': 1,
    'react/no-did-update-set-state': 1,
    'react/no-unknown-property': 1,
    'react/prop-types': 0,
    'react/self-closing-comp': 0,
    'react/sort-comp': 0,
    'react/jsx-wrap-multilines': 1,
    // allowExpressions allow single expressions in a fragment eg: <>{children}</>
    // https://github.com/jsx-eslint/eslint-plugin-react/blob/f83b38869c7fc2c6a84ef8c2639ac190b8fef74f/docs/rules/jsx-no-useless-fragment.md#allowexpressions
    'react/jsx-no-useless-fragment': [2, { allowExpressions: true }],

    'react-hooks/rules-of-hooks': 1,
    'react-hooks/exhaustive-deps': 1,

    // Turned off because we use automatic runtime.
    'react/jsx-uses-react': 0,
    'react/react-in-jsx-scope': 0,
  },
  settings: {
    react: {
      // React version. "detect" automatically picks the version you have installed.
      version: 'detect',
    },
    'import/resolver': {
      typescript: {},
    },
  },
};
