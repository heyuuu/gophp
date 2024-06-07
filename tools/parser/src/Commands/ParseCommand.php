<?php

namespace GoPhp\Tools\Commands;

use GoPhp\Tools\Parse\Application;
use Symfony\Component\Console\Attribute\AsCommand;
use Symfony\Component\Console\Command\Command;
use Symfony\Component\Console\Input\InputInterface;
use Symfony\Component\Console\Input\InputOption;
use Symfony\Component\Console\Output\OutputInterface;

#[AsCommand(
    name: "parse",
    description: "parse php code to ast data",
)]
class ParseCommand extends Command
{
    protected function configure(): void
    {
        $this->addOption('code', 'c', InputOption::VALUE_REQUIRED, 'php code');
    }

    protected function execute(InputInterface $input, OutputInterface $output): int
    {
        $code = $input->getOption('code');

        $app    = new Application();
        $result = $app->parseCode($code);

        $output->write($result);

        return Command::SUCCESS;
    }
}