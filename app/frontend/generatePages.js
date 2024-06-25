const fs = require('fs');
const path = require('path');

const pagesDir = path.join(__dirname, 'src/pages/admin');
const outputFilePath = path.join(__dirname, 'src/pages/index.ts');

const generatePagesIndex = () => {
  const files = fs.readdirSync(pagesDir);

  const imports = files
    .filter(file => file.endsWith('.tsx'))
    .map((file, index) => {
      const fileName = path.parse(file).name;
      return `import Page${index} from './admin/${fileName}';`;
    });

  const exports = files
    .filter(file => file.endsWith('.tsx'))
    .map((file, index) => {
      const fileName = path.parse(file).name;
      const routePath = fileName.toLowerCase();
      return `{ path: '/${routePath}', component: Page${index} }`;
    });

  const fileContent = `${imports.join('\n')}\n\nexport const routes = [${exports.join(', ')}];\n`;

  fs.writeFileSync(outputFilePath, fileContent);
};

generatePagesIndex();
