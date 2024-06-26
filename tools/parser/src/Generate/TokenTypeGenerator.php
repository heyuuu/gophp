<?php

namespace GoPhp\Tools\Generate;

class TokenTypeGenerator extends BaseGenerator
{
    private string $outputFile   = MAIN_ROOT . '/compile/token/token.go';
    private array  $tokens       = [];
    private array  $tokenSamples = [];
    private string $template     = <<<'TEMPLATE'
+ - * / ++ -- %
TEMPLATE;

    public function generate()
    {
        $this->initTokens();

        $defines = "";
        foreach ($this->tokens as $name => $value) {
            $newName = $this->newTokenName($name);
            $defines .= "    $newName Token = $value\n";
            if (!empty($this->tokenSamples[$name])) {
                $defines .= " // " . $this->tokenSamples[$name];
            }
        }

        $content = <<<CODE
package token

type Token int

const (
    {$defines}
)
CODE;

        $this->writeFile($this->outputFile, $content);
    }

    private function initTokens()
    {
        // init tokens
        $tokens    = [];
        $constants = get_defined_constants(true)['tokenizer'];
        foreach ($constants as $name => $value) {
            if (str_starts_with($name, 'T_')) {
                $tokens[$name] = $value;
            }
        }
        asort($tokens);
        $this->tokens = $tokens;

        // init token samples
        foreach (explode(' ', $this->template) as $code) {
            $tokens = token_get_all('<?php ' . $code);
            // var_dump($tokens);
        }
    }

    private function newTokenName(string $name): string
    {
        $name = substr($name, 2);
        return $name;
    }
}