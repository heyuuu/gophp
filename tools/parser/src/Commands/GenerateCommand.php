<?php

namespace GoPhp\Tools\Commands;

use GoPhp\Tools\Generate\AstDecodeGenerator;
use GoPhp\Tools\Generate\AstNodeGenerator;
use GoPhp\Tools\Generate\TokenTypeGenerator;
use Symfony\Component\Console\Attribute\AsCommand;
use Symfony\Component\Console\Command\Command;
use Symfony\Component\Console\Input\InputInterface;
use Symfony\Component\Console\Input\InputOption;
use Symfony\Component\Console\Output\OutputInterface;

#[AsCommand(
    name: "generate",
    description: "generate go code",
)]
class GenerateCommand extends Command
{
    protected function configure()
    {
        $this->addOption('mode', 'm', InputOption::VALUE_REQUIRED, 'generate mode');
    }

    protected function execute(InputInterface $input, OutputInterface $output): int
    {
        $mode = $input->getOption('mode');

        $generator = match ($mode) {
            'ast-node' => new AstNodeGenerator(),
            'ast-decode' => new AstDecodeGenerator(),
            'token-type' => new TokenTypeGenerator(),
            default => throw new \RuntimeException("不支持的 generate mode: {$mode}"),
        };

        $generator->generate();

        return Command::SUCCESS;
    }
}